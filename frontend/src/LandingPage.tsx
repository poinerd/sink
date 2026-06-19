import React from 'react';

export const LandingPage: React.FC = () => {
  return (
    <div className="min-h-screen bg-neutral-50 text-neutral-900 antialiased font-sans">
      
      {/* Navigation */}
      <nav className="max-w-5xl mx-auto px-6 py-6 flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <span className="text-xl font-bold tracking-tight">Sink</span>
          <span className="text-xs bg-neutral-200 text-neutral-700 font-mono px-2 py-0.5 rounded">v1.0</span>
        </div>
        <div className="flex items-center space-x-6 text-sm font-medium">
          <a href="#features" className="text-neutral-500 hover:text-neutral-900 transition-colors">Features</a>
          <a href="#docs" className="text-neutral-500 hover:text-neutral-900 transition-colors">Documentation</a>
          <button className="bg-neutral-900 hover:bg-neutral-800 text-white px-4 py-1.5 rounded-md transition-colors text-xs">
            Sign In
          </button>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="max-w-4xl mx-auto px-6 pt-20 pb-16 text-center">
        <h1 className="text-5xl font-bold tracking-tight text-neutral-900 sm:text-6xl max-w-2xl mx-auto leading-[1.1]">
          Form endpoints for modern developers.
        </h1>
        <p className="mt-6 text-base text-neutral-500 max-w-xl mx-auto leading-relaxed">
          A high-performance, multi-tenant backend engine built from first principles. 
          Ingest unstructured form payloads instantly into your independent PostgreSQL infrastructure.
        </p>
        <div className="mt-10 flex items-center justify-center space-x-4">
          <button className="bg-neutral-900 hover:bg-neutral-800 text-white px-5 py-2.5 rounded-md text-sm font-medium transition-colors shadow-sm">
            Create Free Account
          </button>
          <a href="#docs" className="border border-neutral-300 hover:bg-neutral-100 text-neutral-700 px-5 py-2.5 rounded-md text-sm font-medium transition-colors">
            Read the Docs
          </a>
        </div>
      </section>

      {/* Code Demo Section */}
      <section id="docs" className="max-w-3xl mx-auto px-6 py-12">
        <div className="bg-neutral-900 rounded-xl border border-neutral-800 shadow-xl overflow-hidden">
          <div className="flex items-center space-x-2 bg-neutral-950 px-4 py-3 border-b border-neutral-800">
            <div className="w-3 h-3 rounded-full bg-neutral-800"></div>
            <div className="w-3 h-3 rounded-full bg-neutral-800"></div>
            <div className="w-3 h-3 rounded-full bg-neutral-800"></div>
            <span className="text-xs text-neutral-500 font-mono pl-2">contact-form.html</span>
          </div>
          <div className="p-6 overflow-x-auto font-mono text-xs leading-relaxed text-neutral-300">
            <pre>
{`<form 
  action="https://sink.dev/api/v1/forms/submit" 
  method="POST"
  headers={{ "X-Tenant-ID": "your_tenant_uuid" }}
>
  <label>Email Address</label>
  <input type="email" name="email" required />

  <label>Message</label>
  <textarea name="message" required></textarea>

  <button type="submit">Submit to Engine</button>
</form>`}
            </pre>
          </div>
        </div>
        <p className="text-center text-xs text-neutral-400 mt-4 font-mono">
          Zero client-side setup required. Point your markup directly to Sink.
        </p>
      </section>

      {/* Feature Grid Section */}
      <section id="features" className="max-w-5xl mx-auto px-6 py-20 border-t border-neutral-200">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          
          <div className="space-y-2">
            <div className="h-8 w-8 bg-neutral-900 text-white rounded-lg flex items-center justify-center font-bold text-sm">1</div>
            <h3 className="font-semibold text-neutral-900 text-base">Strict Multi-Tenancy</h3>
            <p className="text-sm text-neutral-500 leading-relaxed">
              Complete data separation boundaries ensure independent client environments never blend or bleed.
            </p>
          </div>

          <div className="space-y-2">
            <div className="h-8 w-8 bg-neutral-900 text-white rounded-lg flex items-center justify-center font-bold text-sm">2</div>
            <h3 className="font-semibold text-neutral-900 text-base">Dynamic Schema Ingestion</h3>
            <p className="text-sm text-neutral-500 leading-relaxed">
              No migrations needed. Send any raw JSON layout from your frontend, and Sink maps it seamlessly to your target tables.
            </p>
          </div>

          <div className="space-y-2">
            <div className="h-8 w-8 bg-neutral-900 text-white rounded-lg flex items-center justify-center font-bold text-sm">3</div>
            <h3 className="font-semibold text-neutral-900 text-base">Zero Bloat Performance</h3>
            <p className="text-sm text-neutral-500 leading-relaxed">
              Engineered cleanly in Go using native standard library primitives for optimal raw TCP request pipelines and connection processing.
            </p>
          </div>

        </div>
      </section>

      {/* Footer */}
      <footer className="max-w-5xl mx-auto px-6 py-12 border-t border-neutral-200 text-center text-xs text-neutral-400">
        <p>&copy; {new Date().getFullYear()} Sink Engine Project. Built from first principles.</p>
      </footer>

    </div>
  );
};