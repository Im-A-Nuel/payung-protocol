"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useConfig, useWriteContract } from "wagmi";
import { waitForTransactionReceipt } from "wagmi/actions";

import { claimFaucet, getWallet, type Zona } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { IDRP_ADDRESS, POOL_ADDRESS, idrpAbi, poolAbi } from "@/lib/contracts";
import { formatRupiah, isLessThan, multiplyWei, toWei } from "@/lib/format";
import { pesanGalat } from "@/lib/galat";
import { Berhasil, Kartu, Tombol } from "./ui";

const PILIHAN_MINGGU = [1, 2, 3, 4];

export function PanelBeli({ zona, terkunci }: { zona: Zona; terkunci: boolean }) {
  const { alamat, bisaKirim } = useAkun();
  const router = useRouter();
  const queryClient = useQueryClient();
  const wagmiConfig = useConfig();
  const { writeContractAsync } = useWriteContract();

  const [minggu, setMinggu] = useState(1);
  const [galat, setGalat] = useState<string | null>(null);

  const dompet = useQuery({
    queryKey: ["dompet", alamat],
    queryFn: () => getWallet(alamat as string),
    enabled: Boolean(alamat),
  });

  const total = multiplyWei(zona.premiumPerWeek, minggu);
  const saldoKurang = dompet.data ? isLessThan(dompet.data.balance, total) : false;

  const faucet = useMutation({
    mutationFn: () => claimFaucet(alamat as string),
    onSuccess: async () => {
      setGalat(null);
      await queryClient.invalidateQueries({ queryKey: ["dompet", alamat] });
    },
    onError: (e) => setGalat(pesanGalat(e)),
  });

  const beli = useMutation({
    mutationFn: async () => {
      if (!dompet.data) throw new Error("Saldo belum termuat");

      // One tap for the driver; approve only runs when the pool still needs
      // permission to pull the premium.
      if (isLessThan(dompet.data.allowance, total)) {
        const hashIzin = await writeContractAsync({
          address: IDRP_ADDRESS,
          abi: idrpAbi,
          functionName: "approve",
          args: [POOL_ADDRESS, toWei(total)],
        });
        await waitForTransactionReceipt(wagmiConfig, { hash: hashIzin });
      }

      const hashBeli = await writeContractAsync({
        address: POOL_ADDRESS,
        abi: poolAbi,
        functionName: "buyPolicy",
        args: [zona.id, minggu],
      });
      await waitForTransactionReceipt(wagmiConfig, { hash: hashBeli });
    },
    onSuccess: async () => {
      setGalat(null);
      await queryClient.invalidateQueries();
    },
    onError: (e) => setGalat(pesanGalat(e)),
  });

  const sedangProses = beli.isPending;

  // The driver sees the result where they tapped, and moves on when they
  // are ready, rather than being thrown to another screen mid-thought.
  if (beli.isSuccess) {
    return (
      <Kartu className="mt-4" id="panel-beli">
        <Berhasil
          judul="Berhasil, kamu terlindungi"
          pesan={`Polis ${zona.name} mulai jalan besok pagi dan aktif ${minggu * 7} hari. Kalau hujan lewat ${zona.thresholdMm} mm, uangnya masuk sendiri.`}
          anak={
            <Tombol onClick={() => router.push("/polis?baru=1")}>Lihat polis saya</Tombol>
          }
        />
      </Kartu>
    );
  }

  return (
    <Kartu className="mt-4" id="panel-beli">
      <fieldset disabled={sedangProses}>
        <legend className="text-[15px] font-bold">Mau dilindungi berapa minggu?</legend>

        <div className="mt-3 grid grid-cols-4 gap-2" role="group" aria-label="Lama perlindungan">
          {PILIHAN_MINGGU.map((n) => {
            const aktif = n === minggu;
            return (
              <button
                key={n}
                type="button"
                onClick={() => setMinggu(n)}
                aria-pressed={aktif}
                className={`tekan min-h-[52px] rounded-xl border text-[16px] font-bold transition-colors ${
                  aktif ? "border-langit bg-langit text-white" : "border-garis bg-kartu text-tinta"
                }`}
              >
                {n}
              </button>
            );
          })}
        </div>
        <p className="mt-2 text-[13px] text-abu">
          Perlindungan mulai besok pagi dan jalan {minggu * 7} hari.
        </p>

        <div className="mt-4 flex items-end justify-between border-t border-garis pt-4">
          <div>
            <p className="text-[13px] text-abu">Yang kamu bayar</p>
            <p className="angka text-[28px] leading-tight font-extrabold">{formatRupiah(total)}</p>
          </div>
          {dompet.data ? (
            <p className="angka text-right text-[13px] text-abu">
              Saldo kamu
              <br />
              <span className="font-bold text-tinta">{formatRupiah(dompet.data.balance)}</span>
            </p>
          ) : null}
        </div>
      </fieldset>

      {saldoKurang ? (
        <div className="mt-4 rounded-xl bg-kertas p-3">
          <p className="text-[14px] font-semibold">Saldomu belum cukup buat premi ini.</p>
          <p className="mt-0.5 text-[13px] text-abu">
            Ambil saldo uji coba dulu. Gratis, sekali sehari, cuma berlaku di jaringan testnet.
          </p>
          <Tombol
            varian="halus"
            className="mt-3"
            sedangProses={faucet.isPending}
            labelProses="Sedang diproses"
            onClick={() => faucet.mutate()}
            disabled={!alamat || sedangProses}
          >
            Ambil saldo testnet
          </Tombol>
        </div>
      ) : null}

      <Tombol
        className="mt-4"
        sedangProses={sedangProses}
        onClick={() => beli.mutate()}
        disabled={terkunci || saldoKurang || !bisaKirim || dompet.isLoading}
      >
        {terkunci ? "Polismu masih jalan di zona ini" : `Lindungi ${minggu} minggu`}
      </Tombol>

      {!bisaKirim ? (
        <p className="mt-2 text-center text-[13px] text-abu">
          Mode pratinjau. Pembelian dimatikan sampai login diaktifkan.
        </p>
      ) : null}

      {galat ? (
        <p role="alert" className="mt-3 text-center text-[14px] font-semibold text-bahaya">
          {galat}
        </p>
      ) : null}

      {faucet.isSuccess && !galat ? (
        <p className="mt-3 text-center text-[14px] font-semibold text-uang">Saldo testnet berhasil masuk.</p>
      ) : null}
    </Kartu>
  );
}
