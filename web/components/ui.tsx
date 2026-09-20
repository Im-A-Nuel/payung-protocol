import type { ButtonHTMLAttributes, ReactNode } from "react";

type Varian = "utama" | "kedua" | "halus";

const GAYA: Record<Varian, string> = {
  utama: "bg-langit text-white active:bg-langit-tua disabled:bg-langit/40",
  kedua: "border border-garis bg-kartu text-tinta active:bg-kertas disabled:text-abu",
  halus: "bg-langit-muda text-langit active:bg-langit-muda/70 disabled:text-abu",
};

type TombolProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  varian?: Varian;
  sedangProses?: boolean;
  labelProses?: string;
};

export function Tombol({
  varian = "utama",
  sedangProses = false,
  labelProses = "Sedang diproses",
  children,
  className = "",
  disabled,
  ...props
}: TombolProps) {
  return (
    <button
      {...props}
      disabled={disabled || sedangProses}
      aria-busy={sedangProses || undefined}
      className={`flex min-h-[52px] w-full items-center justify-center gap-2 rounded-2xl px-5 text-[15px] font-bold transition-colors ${GAYA[varian]} ${className}`}
    >
      {sedangProses ? (
        <>
          <Pemuat />
          {labelProses}
        </>
      ) : (
        children
      )}
    </button>
  );
}

function Pemuat() {
  return (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" aria-hidden="true" className="animate-spin">
      <circle cx="12" cy="12" r="9" stroke="currentColor" strokeOpacity="0.3" strokeWidth="3" />
      <path d="M21 12a9 9 0 0 0-9-9" stroke="currentColor" strokeWidth="3" strokeLinecap="round" />
    </svg>
  );
}

export function Kartu({ children, className = "" }: { children: ReactNode; className?: string }) {
  return (
    <div className={`rounded-2xl border border-garis bg-kartu p-4 ${className}`}>{children}</div>
  );
}

export function JudulHalaman({ judul, anak }: { judul: string; anak?: string }) {
  return (
    <header className="mb-5">
      <h1 className="text-[26px] leading-tight font-extrabold tracking-tight">{judul}</h1>
      {anak ? <p className="mt-1 text-[15px] text-abu">{anak}</p> : null}
    </header>
  );
}

export function Memuat({ tinggi = "h-24" }: { tinggi?: string }) {
  return (
    <div
      role="status"
      aria-label="Memuat"
      className={`w-full animate-pulse rounded-2xl border border-garis bg-kartu ${tinggi}`}
    />
  );
}

export function Galat({ pesan, onCoba }: { pesan: string; onCoba?: () => void }) {
  return (
    <Kartu className="text-center">
      <p className="text-[15px] font-semibold text-bahaya">{pesan}</p>
      {onCoba ? (
        <button
          onClick={onCoba}
          className="mt-3 min-h-[44px] rounded-xl px-4 text-[15px] font-bold text-langit active:text-langit-tua"
        >
          Coba lagi
        </button>
      ) : null}
    </Kartu>
  );
}

export function Kosong({ judul, pesan, anak }: { judul: string; pesan: string; anak?: ReactNode }) {
  return (
    <Kartu className="text-center">
      <p className="text-[17px] font-bold">{judul}</p>
      <p className="mt-1 text-[15px] leading-relaxed text-abu">{pesan}</p>
      {anak ? <div className="mt-4">{anak}</div> : null}
    </Kartu>
  );
}
