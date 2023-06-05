import { useCallback, useState } from "react";
import { useSenderStore } from "../../store/module/sender";
import Icon from "../Icon";

function SenderManagementForm() {
  const senderStore = useSenderStore();
  const { senders } = senderStore.state;
  const [form, setForm] = useState({
    host: "",
    port: "",
    email: "",
    password: "",
  });

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
      await senderStore.createSender({
        host: form.host,
        port: +form.port,
        email: form.email,
        password: form.password,
      });
      setForm({ host: "", port: "", email: "", password: "" });
    },
    [form, setForm]
  );

  return (
    <form className="flex flex-wrap mb-4" onSubmit={handleSubmitForm}>
      <div className="w-full flex items-center">
        <h2 className="text-xl text-gray-700 font-bold">
          발신자 관리 ({senders.length}개)
        </h2>
        <div className="p-2 flex flex-1 justify-end gap-1">
          <button className="w-12 px-4" type="submit">
            <Icon.PlusSquare className="w-6 h-auto opacity-70" />
          </button>
        </div>
      </div>
      <div className="py-2 w-1/2">
        <input
          name="host"
          value={form.host}
          onChange={handleChangeForm}
          placeholder="SMTP 서버명"
          className="py-2 px-2 w-full text-sm  border-b rounded-sm"
          required
        />
      </div>
      <div className="p-2 w-1/2">
        <input
          name="port"
          value={form.port}
          onChange={handleChangeForm}
          type="number"
          placeholder="포트 정보"
          className="py-2 px-2 w-full text-sm  border-b rounded-sm"
          required
        />
      </div>
      <div className="py-2 w-1/2">
        <input
          name="email"
          value={form.email}
          onChange={handleChangeForm}
          placeholder="이메일"
          className="py-2 px-2 w-full text-sm  border-b rounded-sm"
          required
        />
      </div>
      <div className="p-2 w-1/2">
        <input
          name="password"
          value={form.password}
          onChange={handleChangeForm}
          type="password"
          placeholder="비밀번호"
          className="py-2 px-2 w-full text-sm  border-b rounded-sm"
          required
        />
      </div>
    </form>
  );
}

export default SenderManagementForm;
