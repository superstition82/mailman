import SenderManagementForm from "../components/sender-management/SenderManagementForm";
import SenderManagementTable from "../components/sender-management/SenderManagementTable";

function SenderDashboard() {
  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <SenderManagementForm />
        <SenderManagementTable />
      </div>
    </div>
  );
}

export default SenderDashboard;
