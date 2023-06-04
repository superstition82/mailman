import { forwardRef, useMemo, useRef } from "react";
import ReactQuill, { Quill } from "react-quill";
import ImageResize from "quill-image-resize";
import ImageUploader from "quill-image-uploader";
import "react-quill/dist/quill.snow.css";

type Props = {
  className: string;
  initialContent: string;
  placeholder: string;
};
type Ref = ReactQuill;

// https://github.com/quilljs/quill/issues/3077
const Image = Quill.import("formats/image");
Image.sanitize = (url: string) => url;

const DirectionAttribute = Quill.import("attributors/attribute/direction");
const AlignClass = Quill.import("attributors/class/align");
const BackgroundClass = Quill.import("attributors/class/background");
const ColorClass = Quill.import("attributors/class/color");
const DirectionClass = Quill.import("attributors/class/direction");
const FontClass = Quill.import("attributors/class/font");
const SizeClass = Quill.import("attributors/class/size");
const AlignStyle = Quill.import("attributors/style/align");
const BackgroundStyle = Quill.import("attributors/style/background");
const ColorStyle = Quill.import("attributors/style/color");
const DirectionStyle = Quill.import("attributors/style/direction");
const FontStyle = Quill.import("attributors/style/font");
const SizeStyle = Quill.import("attributors/style/size");

Quill.register("modules/ImageResize", ImageResize);
Quill.register("modules/imageUploader", ImageUploader);
Quill.register(DirectionAttribute, true);
Quill.register(AlignClass, true);
Quill.register(BackgroundClass, true);
Quill.register(ColorClass, true);
Quill.register(DirectionClass, true);
Quill.register(FontClass, true);
Quill.register(SizeClass, true);
Quill.register(AlignStyle, true);
Quill.register(BackgroundStyle, true);
Quill.register(ColorStyle, true);
Quill.register(DirectionStyle, true);
Quill.register(FontStyle, true);
Quill.register(SizeStyle, true);

export const Editor = forwardRef<Ref, Props>(function Editor(props, ref) {
  const { className, initialContent, placeholder } = props;
  const editorRef = useRef<HTMLTextAreaElement>(null);

  const modules = useMemo(() => {
    return {
      toolbar: {
        container: [
          ["bold", "italic", "underline", "strike", "blockquote"],
          [{ size: ["small", false, "large", "huge"] }, { color: [] }],
          [
            { list: "ordered" },
            { list: "bullet" },
            { indent: "-1" },
            { indent: "+1" },
            { align: [] },
          ],
          ["link", "image"],
          ["clean"],
        ],
      },
    };
  }, []);

  const formats = useMemo(
    () => [
      "header",
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
      value={initialContent}
      placeholder={placeholder}
      modules={modules}
      formats={formats}
    />
  );
});
