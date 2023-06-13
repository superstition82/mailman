import { useRef } from "react";

type Props = {
  Icon: React.ReactNode;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

function FilePicker({ Icon, onChange }: Props) {
  const fileInput = useRef<HTMLInputElement>(null);

  return (
    <>
      <button color="primary" onClick={() => fileInput.current?.click()}>
        {Icon}
      </button>
      <input ref={fileInput} type="file" onChange={onChange} hidden />
    </>
  );
}

export default FilePicker;
