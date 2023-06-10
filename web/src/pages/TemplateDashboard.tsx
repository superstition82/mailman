import TemplateManagementTable from "../components/template-management/TemplateManagementTable";

function TemplateDashboard() {
  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 pb-8 bg-zinc-100">
      <div className="w-full relative px-4 py-2 rounded-xl bg-white">
        <TemplateManagementTable />
      </div>
    </section>
  );
}

export default TemplateDashboard;
