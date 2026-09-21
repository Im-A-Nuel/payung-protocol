import type { ButtonHTMLAttributes, CSSProperties, ReactNode } from "react";

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
      className={`tekan flex min-h-[52px] w-full items-center justify-center gap-2 rounded-2xl px-5 text-[15px] font-bold transition-colors ${GAYA[varian]} ${className}`}
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

export function Kartu({
  children,
  className = "",
  id,
  style,
}: {
  children: ReactNode;
  className?: string;
  id?: string;
  style?: CSSProperties;
}) {
  return (
    <div id={id} style={style} className={`rounded-2xl border border-garis bg-kartu p-4 ${className}`}>
      {children}
    </div>
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

/**
 * The confirmation a driver gets after something worked. CLAUDE.md is
 * explicit that the word they see is "berhasil", never a transaction hash.
 */
export function Berhasil({ judul, pesan, anak }: { judul: string; pesan?: string; anak?: ReactNode }) {
  return (
    <div className="py-2 text-center">
      <span className="lingkar mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-uang-muda">
        <svg width="34" height="34" viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            className="centang"
            d="M5 12.5 10 17.5 19 7.5"
            stroke="var(--color-uang)"
            strokeWidth="2.6"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </span>

      <p className="masuk mt-3 text-[20px] font-extrabold" style={{ "--tunda": "240ms" } as CSSProperties}>
        {judul}
      </p>
      {pesan ? (
        <p
          className="masuk mt-1 text-[15px] leading-relaxed text-abu"
          style={{ "--tunda": "320ms" } as CSSProperties}
        >
          {pesan}
        </p>
      ) : null}
      {anak ? (
        <div className="masuk mt-5" style={{ "--tunda": "400ms" } as CSSProperties}>
          {anak}
        </div>
      ) : null}
    </div>
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
