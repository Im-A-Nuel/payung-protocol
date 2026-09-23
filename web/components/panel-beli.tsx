"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useBalance, useConfig, useWriteContract } from "wagmi";
import { waitForTransactionReceipt } from "wagmi/actions";
import { formatEther } from "viem";

import { claimFaucet, getWallet, type Zona } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { IDRP_ADDRESS, POOL_ADDRESS, idrpAbi, poolAbi, rantai } from "@/lib/contracts";
import { formatRupiah, isLessThan, multiplyWei, toWei } from "@/lib/format";
import { pesanGalat } from "@/lib/galat";
import { Galat, Kartu, Tombol } from "./ui";

const PILIHAN_MINGGU = [1, 2, 3, 4];

export function PanelBeli({
  zona,
  terkunci,
  tampilkanNarasi,
}: {
  zona: Zona;
  terkunci: boolean;
  tampilkanNarasi: boolean;
}) {
  const { alamat, bisaKirim, bisaMasuk } = useAkun();
  const router = useRouter();
  const queryClient = useQueryClient();
  const wagmiConfig = useConfig();
  const { writeContractAsync } = useWriteContract();

  const [minggu, setMinggu] = useState(1);
  const [galat, setGalat] = useState<string | null>(null);
  const [tahap, setTahap] = useState<"izin" | "beli" | null>(null);
  const [alamatDisalin, setAlamatDisalin] = useState(false);

  const dompet = useQuery({
    queryKey: ["dompet", alamat],
    queryFn: () => getWallet(alamat as string),
    enabled: Boolean(alamat),
  });

  const total = multiplyWei(zona.premiumPerWeek, minggu);
  const premiTersedia = toWei(zona.premiumPerWeek) > 0n;
  const saldoKurang = dompet.data ? isLessThan(dompet.data.balance, total) : false;
  const gas = useBalance({
    address: alamat ?? undefined,
    chainId: rantai.id,
    query: { enabled: Boolean(alamat && bisaKirim) },
  });
  const gasKosong = gas.data?.value === 0n;
  const gasSiap = gas.isSuccess && gas.data.value > 0n;

  async function salinAlamat() {
    if (!alamat) return;
    try {
      await navigator.clipboard.writeText(alamat);
      setAlamatDisalin(true);
    } catch {
      setGalat("Alamat belum bisa disalin otomatis. Tekan lama alamat di atas untuk menyalinnya.");
    }
  }

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
      setTahap(null);

      if (isLessThan(dompet.data.allowance, total)) {
        setTahap("izin");
        const hashIzin = await writeContractAsync({
          address: IDRP_ADDRESS,
          abi: idrpAbi,
          functionName: "approve",
          args: [POOL_ADDRESS, toWei(total)],
        });
        await waitForTransactionReceipt(wagmiConfig, { hash: hashIzin });
      }

      setTahap("beli");
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
      setTahap(null);
      await queryClient.invalidateQueries();
      router.push("/polis?baru=1");
    },
    onError: (e) => {
      setTahap(null);
      setGalat(pesanGalat(e));
    },
  });

  const sedangProses = beli.isPending;

  return (
    <Kartu className="mt-5 p-5">
      <p className="text-[13px] font-semibold text-langit">Zona pilihanmu</p>
      <h2 className="mt-1 text-[21px] font-extrabold">{zona.name}</h2>
      {tampilkanNarasi && zona.premiumNarrative ? (
        <p className="mt-1 text-[13px] leading-relaxed text-abu">{zona.premiumNarrative}</p>
      ) : null}
      <fieldset disabled={sedangProses}>
        <legend className="mt-5 text-[15px] font-bold">Mau dilindungi berapa minggu?</legend>

        <div className="mt-3 grid grid-cols-4 gap-2" role="group" aria-label="Lama perlindungan">
          {PILIHAN_MINGGU.map((n) => {
            const aktif = n === minggu;
            return (
              <button
                key={n}
                type="button"
                onClick={() => setMinggu(n)}
                aria-pressed={aktif}
                className={`min-h-[52px] rounded-xl border text-[16px] font-bold transition-colors ${
                  aktif ? "border-langit bg-langit text-white" : "border-garis bg-kartu text-tinta"
                }`}
              >
                {n}
              </button>
            );
          })}
        </div>
        <p className="mt-2 text-[13px] text-abu">
          Mulai besok, berlaku {minggu * 7} hari.
        </p>

        <div className="mt-4 flex items-end justify-between border-t border-garis pt-4">
          <div>
            <p className="text-[13px] text-abu">Yang kamu bayar</p>
            <p className="angka text-[28px] leading-tight font-extrabold">{premiTersedia ? formatRupiah(total) : "Belum tersedia"}</p>
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

      {dompet.isError ? (
        <div className="mt-4">
          <Galat pesan="Saldo IDRP belum bisa diperiksa." onCoba={() => dompet.refetch()} />
        </div>
      ) : null}

      {gasKosong ? (
        <div className="mt-4 rounded-xl border border-[#e8b7b7] bg-[#fff1f1] p-4">
          <p className="text-[14px] font-bold text-[#912626]">Dompet ini belum punya tBNB</p>
          <p className="mt-1 text-[13px] leading-relaxed text-tinta">
            Kamu perlu sedikit tBNB di opBNB Testnet untuk menyetujui IDRP dan membeli polis.
            Kirim ke alamat dompet ini, lalu cek saldo lagi.
          </p>
          {alamat ? (
            <>
              <p className="mt-3 break-all rounded-lg bg-white px-3 py-2 font-mono text-[12px] text-tinta">
                {alamat}
              </p>
              <button
                type="button"
                onClick={salinAlamat}
                className="mt-2 min-h-[44px] rounded-lg px-2 text-[14px] font-bold text-langit"
              >
                {alamatDisalin ? "Alamat tersalin" : "Salin alamat dompet"}
              </button>
            </>
          ) : null}
          <button type="button" onClick={() => gas.refetch()} className="ml-2 min-h-[44px] rounded-lg px-2 text-[14px] font-bold text-langit">
            Cek saldo lagi
          </button>
        </div>
      ) : null}

      {gas.isError ? (
        <div className="mt-4 rounded-xl bg-kertas p-3">
          <p className="text-[13px] leading-relaxed text-abu">
            Saldo tBNB belum bisa diperiksa. Periksa koneksi jaringan, lalu coba lagi.
          </p>
          <button
            type="button"
            onClick={() => gas.refetch()}
            className="mt-1 min-h-[44px] rounded-lg px-2 text-[14px] font-bold text-langit"
          >
            Periksa lagi
          </button>
        </div>
      ) : null}

      {bisaKirim && gas.isPending ? (
        <p role="status" className="mt-4 text-[13px] text-abu">
          Memeriksa saldo tBNB untuk biaya jaringan...
        </p>
      ) : null}

      {gas.data && !gasKosong ? (
        <p className="mt-4 text-[13px] text-abu">
          Saldo gas:{" "}
          <span className="angka font-bold text-tinta">
            {gas.data.value < 1_000_000_000_000n
              ? "<0,000001"
              : Number(formatEther(gas.data.value)).toLocaleString("id-ID", { maximumFractionDigits: 6 })}{" "}
            tBNB
          </span>
        </p>
      ) : null}

      <Tombol
        className="mt-4"
        sedangProses={sedangProses}
        labelProses={tahap === "izin" ? "Menyetujui IDRP" : "Membeli polis"}
        onClick={() => beli.mutate()}
        disabled={terkunci || !premiTersedia || saldoKurang || !gasSiap || !bisaKirim || !dompet.data || dompet.isLoading}
      >
        {terkunci ? "Polismu masih jalan di zona ini" : premiTersedia ? `Lindungi ${minggu} minggu` : "Premi zona belum tersedia"}
      </Tombol>

      {!bisaKirim ? (
        <p className="mt-2 text-center text-[13px] text-abu">
          {bisaMasuk
            ? "Dompet sedang disambungkan. Tunggu sebentar, lalu coba lagi."
            : "Mode pratinjau. Pembelian belum aktif di lingkungan ini."}
        </p>
      ) : null}

      {bisaKirim && !terkunci && !saldoKurang && !gasKosong ? (
        <p className="mt-2 text-center text-[12px] leading-relaxed text-abu">
          Pembelian pertama bisa meminta dua persetujuan dompet: izin IDRP dan beli polis.
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
