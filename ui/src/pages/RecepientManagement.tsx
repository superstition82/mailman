import RecepientManagementForm from "../components/recepient-management/RecepientManagementForm";
import RecepientManagementTable from "../components/recepient-management/RecepientManagementTable";

function RecepientManagement() {
  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <RecepientManagementForm />
        <RecepientManagementTable />
      </div>
    </div>
  );
}

export default RecepientManagement;
