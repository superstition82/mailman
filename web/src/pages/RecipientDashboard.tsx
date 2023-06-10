import RecipientManagementForm from "../components/recipient-management/RecipientManagementForm";
import RecipientManagementTable from "../components/recipient-management/RecipientManagementTable";

function RecipientDashboard() {
  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 pb-8 bg-zinc-100">
      <div className="w-full relative">
        <RecipientManagementForm />
        <RecipientManagementTable />
      </div>
    </section>
  );
}

export default RecipientDashboard;
