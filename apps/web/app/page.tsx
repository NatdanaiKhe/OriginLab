export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-8 bg-paper text-ink">
      <div className="w-full max-w-xl p-10 bg-paper-soft border-2 border-ink rounded-[16px] shadow-[4px_4px_0px_0px_var(--ink)] space-y-6">
        <div className="inline-block px-3 py-1 text-xs font-semibold uppercase tracking-wider bg-paper-deep border border-ink rounded-full">
          OriginLab
        </div>
        <h1 className="text-4xl sm:text-5xl font-semibold tracking-tight leading-[1.05]">
          Roasted small, shipped fresh.
        </h1>
        <p className="text-lg text-ink-soft leading-relaxed">
          Single-origin specialty coffee roaster. Monorepo architecture initialized with modular backend boundaries.
        </p>
        <div className="pt-2 flex gap-4">
          <a
            href="http://localhost:8080/healthz"
            target="_blank"
            rel="noreferrer"
            className="inline-flex items-center justify-center px-5 py-2.5 text-sm font-medium bg-ink text-paper border-2 border-ink rounded-[12px] hover:-translate-y-0.5 hover:shadow-[4px_4px_0px_0px_var(--ink)] transition-all duration-150"
          >
            Check API Health
          </a>
        </div>
      </div>
    </main>
  );
}
