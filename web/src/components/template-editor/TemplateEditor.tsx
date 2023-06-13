import ReactQuill from "react-quill";
import { useEffect, useRef, useState } from "react";
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
  const editorRef = useRef<ReactQuill>(null);
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
            setResourceIdList(template.resourceIdList ?? []);
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
    // https://github.com/quilljs/quill/issues/1328
    const processedHtml = body.replace("<p><br></p>", "").trim();
    if (templateId) {
      await templateStore.patchTemplate({
        id: templateId,
        subject,
        body: processedHtml,
        resourceIdList,
      });
    } else {
      await templateStore.createTemplate({
        subject,
        body: processedHtml,
        resourceIdList,
      });
    }
    navigate("/"); // home
  };

  const handleSubjectInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSubject(e.target.value);
  };

  return (
    <div
      className={`${className} w-full max-w-3xl relative px-4 py-2 rounded-xl bg-white`}
    >
      <div className="editor-container px-4">
        <input
          type="text"
          placeholder="제목"
          className="title-input"
          value={subject}
          onChange={handleSubjectInputChange}
        />
        <div className="editor-wrapper">
          <Editor
            ref={editorRef}
            body={body}
            onChangeBody={setBody}
            onUpload={handleUploadImage}
          />
        </div>
        <WriteActionButtons
          onPublish={handleOnPublish}
          onCancel={() => navigate(-1)}
        />
      </div>
    </div>
  );
}

export default TemplateEditor;
