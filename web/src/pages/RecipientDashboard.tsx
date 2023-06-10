import RecipientManagementForm from "../components/recipient-management/RecipientManagementForm";
import RecipientManagementTable from "../components/recipient-management/RecipientManagementTable";

function RecipientDashboard() {
  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <RecipientManagementForm />
        <RecipientManagementTable />
      </div>
    </div>
  );
}

export default RecipientDashboard;
