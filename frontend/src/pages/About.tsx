export default function About() {
  return (
    <div>
      <div className="page-heading">
        <div>
          <p>BirrWise</p>
          <h1>About</h1>
        </div>
      </div>
      <section className="panel about-card">
        <div className="brand-mark large">B</div>
        <h2>BirrWise</h2>
        <p>
          A modern personal finance tracker designed for everyday life in
          Ethiopia.
        </p>
        <div className="about-grid">
          <div>
            <strong>Local-first</strong>
            <span>Your financial data stays in this browser.</span>
          </div>
          <div>
            <strong>ETB ready</strong>
            <span>Built around the Ethiopian Birr.</span>
          </div>
          <div>
            <strong>Private</strong>
            <span>No account or cloud connection required.</span>
          </div>
        </div>
        <small>Frontend MVP · Version 0.1.0</small>
      </section>
    </div>
  );
}
