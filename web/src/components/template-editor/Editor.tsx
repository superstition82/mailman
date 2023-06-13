import { forwardRef, useMemo } from "react";
import ReactQuill, { Quill } from "react-quill";
import ImageUploader from "quill-image-uploader";
import ImageResize from "quill-image-resize";
import "react-quill/dist/quill.snow.css";
import "../../less/editor.less";

// image format for retrieving custom attributes
const Parchment = Quill.import("parchment");

const BaseImageFormat = Quill.import("formats/image");
const ImageFormatAttributesList = ["alt", "height", "width", "style"];

class ImageFormat extends BaseImageFormat {
  static formats(domNode) {
    return ImageFormatAttributesList.reduce(function (formats, attribute) {
      if (domNode.hasAttribute(attribute)) {
        formats[attribute] = domNode.getAttribute(attribute);
      }
      return formats;
    }, {});
  }
  format(name, value) {
    if (ImageFormatAttributesList.indexOf(name) > -1) {
      if (value) {
        this.domNode.setAttribute(name, value);
      } else {
        this.domNode.removeAttribute(name);
      }
    } else {
      super.format(name, value);
    }
  }
}

Quill.register(ImageFormat, true);

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
        parchment: Parchment,
      },
      imageUploader: {
        upload: handleUpload,
      },
    }),
    []
  );

  return (
    <ReactQuill
      ref={ref}
      theme="snow"
      modules={modules}
      value={body}
      onChange={handleChangeBody}
    />
  );
});

export default Editor;
