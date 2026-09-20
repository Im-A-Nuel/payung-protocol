import { Kartu } from "@/components/ui";

export const metadata = { title: "Tidak ada koneksi" };

export default function HalamanOffline() {
  return (
    <div className="pt-10">
      <Kartu className="text-center">
        <p className="text-[19px] font-extrabold">Sinyalmu putus</p>
        <p className="mt-2 text-[15px] leading-relaxed text-abu">
          Polis dan bayaranmu tetap aman di blockchain. Buka lagi halaman ini begitu sinyal balik.
        </p>
      </Kartu>
    </div>
  );
}
