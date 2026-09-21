"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

const TAB = [
  { href: "/", label: "Lindungi", ikon: IkonPayung },
  { href: "/polis", label: "Polis", ikon: IkonKartu },
  { href: "/riwayat", label: "Riwayat", ikon: IkonRiwayat },
];

export function Nav() {
  const path = usePathname();

  return (
    <nav
      aria-label="Menu utama"
      className="fixed inset-x-0 bottom-0 z-20 border-t border-garis bg-kartu"
      style={{ paddingBottom: "env(safe-area-inset-bottom, 0px)" }}
    >
      <ul className="mx-auto flex w-full max-w-[520px]">
        {TAB.map(({ href, label, ikon: Ikon }) => {
          const aktif = path === href;
          return (
            <li key={href} className="flex-1">
              <Link
                href={href}
                aria-current={aktif ? "page" : undefined}
                className={`tekan relative flex min-h-[60px] flex-col items-center justify-center gap-1 text-[13px] font-semibold transition-colors ${
                  aktif ? "text-langit" : "text-abu"
                }`}
              >
                {aktif ? (
                  <span
                    aria-hidden="true"
                    className="absolute top-0 h-[3px] w-8 rounded-b-full bg-langit"
                  />
                ) : null}
                <Ikon aktif={aktif} />
                {label}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}

function IkonPayung({ aktif }: { aktif: boolean }) {
  return (
    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path
        d="M12 3v1.2M3.5 12.5a8.5 8.5 0 0 1 17 0Z"
        stroke="currentColor"
        strokeWidth={aktif ? 2.2 : 1.8}
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M12 12.5v6.2a2.1 2.1 0 0 1-4.2 0"
        stroke="currentColor"
        strokeWidth={aktif ? 2.2 : 1.8}
        strokeLinecap="round"
      />
    </svg>
  );
}

function IkonKartu({ aktif }: { aktif: boolean }) {
  return (
    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <rect
        x="3"
        y="5.5"
        width="18"
        height="13"
        rx="2.5"
        stroke="currentColor"
        strokeWidth={aktif ? 2.2 : 1.8}
      />
      <path d="M3 10h18" stroke="currentColor" strokeWidth={aktif ? 2.2 : 1.8} strokeLinecap="round" />
      <path d="M7 14.5h4" stroke="currentColor" strokeWidth={aktif ? 2.2 : 1.8} strokeLinecap="round" />
    </svg>
  );
}

function IkonRiwayat({ aktif }: { aktif: boolean }) {
  return (
    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path
        d="M4 6.5h16M4 12h16M4 17.5h10"
        stroke="currentColor"
        strokeWidth={aktif ? 2.2 : 1.8}
        strokeLinecap="round"
      />
    </svg>
  );
}
