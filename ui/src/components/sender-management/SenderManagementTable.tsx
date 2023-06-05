import { useEffect } from "react";
import { Checkbox, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useSenderStore } from "../../store/module/sender";
import Icon from "../Icon";

function SenderManagementTable() {
  const senderStore = useSenderStore();
  const { senders, selectedIds } = senderStore.state;

  useEffect(() => {
    senderStore.fetchSenders();
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    senderStore.toggleSelectAll(event.target.checked);
  };

  const handleClick = (_: React.MouseEvent, id: number) => {
    senderStore.toggleSelect(id);
  };

  const handleDeleteSelected = async () => {
    await senderStore.deleteBulkSender(selectedIds);
  };

  return (
    <div className="flex flex-col rounded-sm px-2 mb-8 bg-white ">
      <Sheet
        sx={{
          "--TableCell-height": "40px",
          "--TableHeader-height": "calc(1 * var(--TableCell-height))",
        }}
      >
        <TableVirtuoso
          style={{ height: "68vh" }}
          data={senders}
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
                    selectedIds.length < senders.length
                  }
                  checked={
                    senders.length > 0 && selectedIds.length === senders.length
                  }
                  onChange={handleSelectAll}
                  sx={{ verticalAlign: "sub" }}
                />
              </th>
              <th>#</th>
              <th>호스트</th>
              <th>이메일</th>
            </tr>
          )}
          itemContent={(_, sender) => (
            <>
              <td onClick={(event) => handleClick(event, sender.id)}>
                <Checkbox
                  checked={selectedIds.includes(sender.id)}
                  sx={{ verticalAlign: "top" }}
                />
              </td>
              <td>{sender.id}</td>
              <td>{sender.host}</td>
              <td>{sender.email}</td>
            </>
          )}
        />
      </Sheet>
      <div className="py-4 flex justify-end gap-1">
        <button className="flex gap-2 px-4" onClick={handleDeleteSelected}>
          <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 선택 삭제
        </button>
      </div>
    </div>
  );
}

export default SenderManagementTable;
