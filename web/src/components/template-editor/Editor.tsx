import { useEffect, useRef } from "react";
import Quill from "quill";
import "quill/dist/quill.bubble.css";
import "../../less/editor.less";

type Props = {
  title: string;
  body: string;
  onChangeTitle: (value: string) => void;
  onChangeBody: (value: string) => void;
};

function Editor({ title, body, onChangeTitle, onChangeBody }: Props) {
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
      placeholder: "본문을 작성하세요...",
      modules: {
        toolbar: [
          [{ header: "1" }, { header: "1" }],
          ["bold", "italic", "underline", "strike"],
          [{ list: "ordered" }, { list: "bullet" }],
          ["blockquote", "code-block", "link", "image"],
        ],
      },
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
