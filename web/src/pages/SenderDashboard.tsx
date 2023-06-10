import SenderManagementForm from "../components/sender-management/SenderManagementForm";
import SenderManagementTable from "../components/sender-management/SenderManagementTable";

function SenderDashboard() {
  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 pb-8 bg-zinc-100">
      <div className="w-full relative px-4 py-2 rounded-xl bg-white">
        <SenderManagementForm />
        <SenderManagementTable />
      </div>
    </section>
  );
}

export default SenderDashboard;
