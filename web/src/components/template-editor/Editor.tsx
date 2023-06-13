import { forwardRef, useMemo } from "react";
import ReactQuill, { Quill } from "react-quill";
import ImageUploader from "quill-image-uploader";
import ImageResize from "quill-image-resize";
import "react-quill/dist/quill.snow.css";
import "../../less/editor.less";

Quill.register("modules/imageUploader", ImageUploader);
Quill.register("modules/ImageResize", ImageResize);

type Props = {
  body: string;
  onChangeBody: (value: string) => void;
  onUpload: (file: File) => Promise<Resource>;
};

type Ref = ReactQuill;

const Editor = forwardRef<Ref, Props>(function Editor(props, ref) {
  const { body, onChangeBody, onUpload } = props;

  const handleUpload = (file: File) => {
    return new Promise((resolve, reject) => {
      onUpload(file)
        .then((resource) => {
          resolve(`/o/r/${resource.id}/${resource.filename}`);
        })
        .catch((error) => {
          console.error("Error:", error);
          reject("Upload failed");
        });
    });
  };

  const handleChangeBody = (value: string) => {
    onChangeBody(value);
  };

  const modules = useMemo(
    () => ({
      toolbar: {
        container: [
          [{ size: ["small", false, "large", "huge"] }, { color: [] }],
          ["bold", "italic", "underline", "strike"],
          [
            { align: "" },
            { align: "center" },
            { align: "right" },
            { align: "justify" },
          ],
          ["blockquote", "code-block", "link", "image"],
          ["bold", "italic", "underline", "strike", "blockquote"],
        ],
      },
      ImageResize: {
        parchment: Quill.import("parchment"),
      },
      imageUploader: {
        upload: handleUpload,
      },
    }),
    []
  );

  const formats = useMemo(
    () => [
      "header",
      "alt",
      "height",
      "width",
      "font",
      "size",
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
      "color",
      "size",
      "video",
      "align",
      "background",
      "direction",
      "code-block",
      "code",
    ],
    []
  );

  return (
    <ReactQuill
      ref={ref}
      theme="snow"
      modules={modules}
      formats={formats}
      value={body}
      onChange={handleChangeBody}
    />
  );
});

export default Editor;
