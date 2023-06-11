import { useCallback, useState } from "react";
import { useRecipientStore } from "../../store/module/recipient";
import Icon from "../common/Icon";

function RecipientManagementForm() {
  const recipientStore = useRecipientStore();
  const { recipients } = recipientStore.state;
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

  const handleSubmitForm = useCallback(
    async (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      await recipientStore.createRecipient({
        email: form.email,
      });
      setForm({ email: "" });
    },
    [form, setForm]
  );

  return (
    <form className="flex flex-wrap mb-4" onSubmit={handleSubmitForm}>
      <div className="w-full flex items-center">
        <p className="flex flex-row justify-start items-center select-none rounded pt-2">
          <Icon.UserPlus className="mr-3 w-6 h-auto opacity-70" />
          수신자 관리 ({recipients.length} 계정)
        </p>
      </div>
      <div className="w-full flex mt-4 bg-white rounded-sm">
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
      </div>
    </form>
  );
}

export default RecipientManagementForm;
