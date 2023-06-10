import TemplateManagementTable from "../components/template-management/TemplateManagementTable";

function TemplateManagement() {
  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <TemplateManagementTable />
      </div>
    </div>
  );
}

export default TemplateManagement;
