import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import WriteActionButtons from "./WriteActionButtons";
import { useTemplateStore } from "../../store/module/template";
import Editor from "./Editor";

type Props = {
  templateId: TemplateId | null;
  className?: string;
  onConfirm?: () => void;
};

function TemplateEditor({ templateId, className, onConfirm }: Props) {
  const navigate = useNavigate();
  const templateStore = useTemplateStore();

  const [subject, setSubject] = useState("");
  const [body, setBody] = useState("");

  useEffect(() => {
    if (templateId) {
      templateStore
        .getTemplateById(templateId)
        .then((template) => {
          if (template) {
            setSubject(template.subject);
            setBody(template.body);
          }
        })
        .catch((error) => {
          console.log(error);
          navigate("/not-found");
        });
    }
  }, [templateId]);

  const handleOnPublish = async () => {
    if (templateId) {
      await templateStore.patchTemplate({
        id: templateId,
        subject,
        body,
      });
    } else {
      await templateStore.createTemplate({
        subject,
        body,
      });
    }
    navigate("/"); // home
  };

  return (
    <div className="w-full max-w-3xl relative px-4 py-2 rounded-xl bg-white">
      <Editor
        title={subject}
        body={body}
        onChangeTitle={setSubject}
        onChangeBody={setBody}
      />
      <WriteActionButtons
        onPublish={handleOnPublish}
        onCancel={() => navigate(-1)}
      />
    </div>
  );
}

export default TemplateEditor;
