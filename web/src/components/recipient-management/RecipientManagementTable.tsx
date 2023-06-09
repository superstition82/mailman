import toast from "react-hot-toast";
import { useCallback, useEffect } from "react";
import { Checkbox, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useRecipientStore } from "../../store/module/recipient";
import Icon from "../common/Icon";
import Progress from "../common/Progress";
import { useLoading } from "../../hooks/useLoading";

function RecipientManagementTable() {
  const recipientStore = useRecipientStore();
  const { recipients, selectedIds } = recipientStore.state;
  const { progress, isLoading, setStart, setNextTick, setFinish } =
    useLoading();

  useEffect(() => {
    recipientStore.fetchRecipients().catch((error) => {
      console.log(error);
      toast.error(error.message);
    });
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    recipientStore.toggleSelectAll(event.target.checked);
  };

  const handleSelect = (id: number) => {
    recipientStore.toggleSelect(id);
  };

  const handleDeleteSelected = async () => {
    await recipientStore.deleteBulkRecipient(selectedIds);
  };

  const handleDeleteUnreachable = async () => {
    const unreachableIds = recipients
      .filter((recipient) => recipient.reachable !== "yes")
      .map((recipient) => recipient.id);

    await recipientStore.deleteBulkRecipient(unreachableIds);
  };

  const handleValidate = useCallback(async () => {
    setStart(selectedIds.length);
    for (const selected of selectedIds) {
      setNextTick();
      await recipientStore.validate(selected);
    }
    setFinish();
  }, [progress, selectedIds, setStart, setNextTick, setFinish]);

  return (
    <div className="flex flex-col rounded-sm px-2 mb-8 bg-white">
      {isLoading ? <Progress value={progress} /> : <></>}
      <Sheet
        sx={{
          "--TableCell-height": "40px",
          "--TableHeader-height": "calc(1 * var(--TableCell-height))",
        }}
      >
        <TableVirtuoso
          style={{ height: "75vh" }}
          data={recipients}
          components={{
            Table: (props) => (
              <Table
                stickyHeader
                hoverRow
                sx={{
                  "& thead th:nth-child(1)": {
                    width: "40px",
                  },
                  "& thead th:nth-child(2)": {
                    width: "80px",
                  },
                }}
                {...props}
              />
            ),
          }}
          fixedHeaderContent={() => (
            <tr>
              <th>
                <Checkbox
                  indeterminate={
                    selectedIds.length > 0 &&
                    selectedIds.length < recipients.length
                  }
                  checked={
                    recipients.length > 0 &&
                    selectedIds.length === recipients.length
                  }
                  onChange={handleSelectAll}
                  sx={{ verticalAlign: "sub" }}
                />
              </th>
              <th>#</th>
              <th>이메일</th>
              <th>전송가능</th>
            </tr>
          )}
          itemContent={(_, recipient) => (
            <>
              <td onClick={() => handleSelect(recipient.id)}>
                <Checkbox
                  checked={selectedIds.includes(recipient.id)}
                  sx={{ verticalAlign: "top" }}
                />
              </td>
              <td>{recipient.id}</td>
              <td>{recipient.email}</td>
              <td>{recipient.reachable}</td>
            </>
          )}
        />
      </Sheet>
      <div className="py-4 flex justify-end gap-1">
        <>
          <button className="flex gap-2 px-4" onClick={handleValidate}>
            <Icon.CheckSquare className="w-6 h-auto opacity-70" /> 선택 검증
          </button>
          <button className="flex gap-2 px-4" onClick={handleDeleteSelected}>
            <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 선택 삭제
          </button>
          <button className="flex gap-2 px-4" onClick={handleDeleteUnreachable}>
            <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 전송불가 삭제
          </button>
        </>
      </div>
    </div>
  );
}

export default RecipientManagementTable;
