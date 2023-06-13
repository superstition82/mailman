import { useCallback, useState } from "react";
import { useRecipientStore } from "../../store/module/recipient";
import Icon from "../common/Icon";
import FilePicker from "../common/FilePicker";
import { useLoading } from "../../hooks/useLoading";
import Progress from "../common/Progress";

function RecipientManagementForm() {
  const recipientStore = useRecipientStore();
  const { recipients } = recipientStore.state;
  const { isLoading, setStart, setFinish } = useLoading();
  const [form, setForm] = useState({ email: "" });

  const handleChangeForm = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setForm((prev) => ({
        ...prev,
        [e.target.name]: e.target.value,
      }));
    },
    [form, setForm]
  );

  const handleSubmitForm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await recipientStore.createRecipient({
      email: form.email,
    });
    setForm({ email: "" });
  };

  const handleUploadFile = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files && e.target.files[0];
    if (!file) return;
    setStart();
    await recipientStore.importRecipientFile(file);
    setFinish();
  };

  return (
    <div className="flex flex-wrap mb-4">
      {isLoading ? <Progress /> : <></>}
      <div className="w-full flex justify-between items-center">
        <p className="flex flex-row justify-start items-center select-none rounded pt-2">
          <Icon.UserPlus className="mr-2 w-6 h-auto opacity-70" />
          수신자 관리 ({recipients.length} 계정)
        </p>
        <div className="flex flex-row justify-center items-center select-none pt-2">
          <FilePicker
            Icon={<Icon.FileUp className="mr-2 w-6 h-auto opacity-70" />}
            onChange={handleUploadFile}
          />
          <a href="/api/recipient/file-export">
            <Icon.FileDown className="mr-2 w-6 h-auto opacity-70" />
          </a>
        </div>
      </div>
      <form
        className="w-full flex mt-4 bg-white rounded-sm"
        onSubmit={handleSubmitForm}
      >
        <input
          className="py-2 px-2 w-full text-sm border rounded"
          value={form.email}
          name="email"
          placeholder="이메일"
          onChange={handleChangeForm}
          required
        />
        <button className="w-12 px-4" type="submit">
          <Icon.PlusSquare className="w-6 h-auto opacity-70" />
        </button>
      </form>
    </div>
  );
}

export default RecipientManagementForm;
