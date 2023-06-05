import { useEffect } from "react";
import { Checkbox, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useSenderStore } from "../../store/module/sender";

function SenderManagementTable() {
  const senderStore = useSenderStore();
  const { senders, selectedIds } = senderStore.state;

  useEffect(() => {
    senderStore.fetchSenders();
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    senderStore.toggleSelectAll(event.target.checked);
  };

  const handleClick = (_: React.MouseEvent<unknown>, id: number) => {
    senderStore.toggleSelect(id);
  };

  const handleDeleteSelected = () => {
    selectedIds.forEach(async (id) => {
      await senderStore.deleteSenderById(id);
    });
  };

  return (
    <div className="flex flex-col rounded-md px-2 mb-8 bg-white ">
      <Sheet
        sx={{
          "--TableCell-height": "40px",
          "--TableHeader-height": "calc(1 * var(--TableCell-height))",
        }}
      >
        <TableVirtuoso
          style={{ height: 400 }}
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
          itemContent={(idx, sender) => (
            <>
              <td onClick={(event) => handleClick(event, sender.id)}>
                <Checkbox
                  checked={selectedIds.includes(sender.id)}
                  sx={{ verticalAlign: "top" }}
                />
              </td>
              <td>{idx + 1}</td>
              <td>{sender.host}</td>
              <td>{sender.email}</td>
            </>
          )}
        />
      </Sheet>
      <div className="p-2 flex justify-end gap-1">
        <button
          className="px-3 py-1 bg-red-500 hover:bg-red-700 text-white text-sm font-bold rounded"
          onClick={handleDeleteSelected}
        >
          선택삭제
        </button>
      </div>
    </div>
  );
}

export default SenderManagementTable;
