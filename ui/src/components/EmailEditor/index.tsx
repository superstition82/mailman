import { useRef } from "react";
import ReactQuill from "react-quill";
import { Editor } from "./Editor";

interface Props {
  className?: string;
  emailId?: string;
  onConfirm?: () => void;
}

interface State {}

export const EmailEditor: React.FC<Props> = (props) => {
  const { className } = props;
  const editorRef = useRef<ReactQuill>(null);

  return (
    <div className={`${className} memo-editor-container`} tabIndex={0}>
      <Editor className="" ref={editorRef} initialContent="" placeholder="" />
    </div>
  );
};
