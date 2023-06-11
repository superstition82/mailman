import { useEffect, useRef } from "react";
import Quill from "quill";
import ImageUploader from "quill-image-uploader";
import ImageResize from "quill-image-resize";
import "quill/dist/quill.bubble.css";
import "quill-image-uploader/dist/quill.imageUploader.min.css";
import "../../less/editor.less";

Quill.register("modules/imageUploader", ImageUploader);
Quill.register("modules/ImageResize", ImageResize);

type Props = {
  title: string;
  body: string;
  onChangeTitle: (value: string) => void;
  onChangeBody: (value: string) => void;
  onUpload: (file: File) => Promise<Resource>;
};

function Editor({ title, body, onChangeTitle, onChangeBody, onUpload }: Props) {
  const quillElement = useRef<HTMLDivElement>(null);
  const quillInstance = useRef<Quill | null>(null);

  const mounted = useRef(false);
  useEffect(() => {
    if (mounted.current || !quillInstance.current) return;
    mounted.current = true;
    quillInstance.current.root.innerHTML = body;
  }, [body]);

  useEffect(() => {
    quillInstance.current = new Quill(quillElement.current!, {
      theme: "bubble",
      formats: [
        "header",
        "bold",
        "italic",
        "underline",
        "strike",
        "blockquote",
        "list",
        "bullet",
        "indent",
        "link",
        "image",
        "imageBlot",
      ],
      modules: {
        toolbar: {
          container: [
            [{ header: "1" }, { header: "2" }],
            ["bold", "italic", "underline", "strike"],
            [{ list: "ordered" }, { list: "bullet" }],
            ["blockquote", "code-block", "link", "image"],
          ],
        },
        imageUploader: {
          upload: (file: File) => {
            return new Promise((resolve, reject) => {
              onUpload(file)
                .then((resource) => {
                  console.log("resource: ", resource);
                  resolve(`/o/r/${resource.id}/${resource.filename}`);
                })
                .catch((error) => {
                  reject("Upload failed");
                  console.error("Error:", error);
                });
            });
          },
        },
        ImageResize: {
          parchment: Quill.import("parchment"),
        },
      },
      placeholder: "본문을 작성하세요...",
    });

    const quill = quillInstance.current;
    quill.on("text-change", (delta, oldDelta, source) => {
      if (source === "user") {
        onChangeBody(quill.root.innerHTML);
      }
    });
  }, [onChangeBody]);

  const handleChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    onChangeTitle(e.target.value);
  };

  return (
    <div className="editor-container px-4">
      <input
        type="text"
        placeholder="제목"
        className="title-input"
        onChange={handleChangeTitle}
        value={title}
      />
      <div className="editor-wrapper">
        <div ref={quillElement} />
      </div>
    </div>
  );
}

export default Editor;
