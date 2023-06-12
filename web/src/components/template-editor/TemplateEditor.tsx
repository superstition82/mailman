import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Editor from "./Editor";
import WriteActionButtons from "./WriteActionButtons";
import { useTemplateStore } from "../../store/module/template";
import { useResourceStore } from "../../store/module/resource";

type Props = {
  templateId: TemplateId | null;
  className?: string;
};

function TemplateEditor({ templateId, className }: Props) {
  const navigate = useNavigate();
  const templateStore = useTemplateStore();
  const resourceStore = useResourceStore();

  const [subject, setSubject] = useState("");
  const [body, setBody] = useState("");
  const [resourceIdList, setResourceIdList] = useState<ResourceId[]>([]);

  useEffect(() => {
    if (templateId) {
      templateStore
        .findTemplateById(templateId)
        .then((template) => {
          if (template) {
            setSubject(template.subject);
            setBody(template.body);
            setResourceIdList(template.resourceIdList);
          }
        })
        .catch((error) => {
          console.log(error);
          navigate("/not-found");
        });
    }
  }, [templateId]);

  const handleUploadImage = async (file: File) => {
    const resource = await resourceStore.createResourceWithBlob(file);
    setResourceIdList((prev) => [...prev, resource.id]);

    return resource;
  };

  const handleOnPublish = async () => {
    if (templateId) {
      await templateStore.patchTemplate({
        id: templateId,
        subject,
        body,
        resourceIdList,
      });
    } else {
      await templateStore.createTemplate({
        subject,
        body,
        resourceIdList,
      });
    }
    navigate("/"); // home
  };

  return (
    <div
      className={`${className} w-full max-w-3xl relative px-4 py-2 rounded-xl bg-white`}
    >
      <Editor
        title={subject}
        body={body}
        onChangeTitle={setSubject}
        onChangeBody={setBody}
        onUpload={handleUploadImage}
      />
      <WriteActionButtons
        onPublish={handleOnPublish}
        onCancel={() => navigate(-1)}
      />
    </div>
  );
}

export default TemplateEditor;
