import _ from "lodash-es";
import { useRecipientStore } from "../store/module/recipient";
import { useSenderStore } from "../store/module/sender";
import { useTemplateStore } from "../store/module/template";

type SendOption = {
  waitMs: number;
  bundle: number;
  bcc: boolean;
};

export const useMailman = () => {
  const templateStore = useTemplateStore(),
    senderStore = useSenderStore(),
    recipientStore = useRecipientStore();

  const { selectedIds: templateIds } = templateStore.state,
    { selectedIds: senderIds } = senderStore.state,
    { selectedIds: recipientIds } = recipientStore.state;

  const send = async (options: SendOption) => {
    for (let senderId of senderIds) {
    }
  };

  return { send };
};
