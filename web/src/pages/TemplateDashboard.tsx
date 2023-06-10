import Icon from "../components/common/Icon";
import TemplateManagementTable from "../components/template-management/TemplateManagementTable";

function TemplateDashboard() {
  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 pb-8 bg-zinc-100">
      <div className="w-full relative px-4 py-2 rounded-xl bg-white">
        <div className="w-full flex items-center mb-4">
          <p className="flex flex-row justify-start items-center select-none rounded pt-2">
            <Icon.Hash className="mr-3 w-6 h-auto opacity-70" />
            이메일
          </p>
        </div>
        <TemplateManagementTable />
      </div>
    </section>
  );
}

export default TemplateDashboard;
