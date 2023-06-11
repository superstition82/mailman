import { useSearchParams } from "react-router-dom";
import TemplateEditor from "../components/template-editor/TemplateEditor";

function TemplateWrite() {
  const [searchParams, _] = useSearchParams();
  const templateId = parseInt(searchParams.get("id") ?? "", 10);

  return (
    <section className="w-full min-h-full flex flex-col justify-start items-center px-4 py-8 bg-zinc-100">
      <TemplateEditor templateId={templateId} />
    </section>
  );
}

export default TemplateWrite;
