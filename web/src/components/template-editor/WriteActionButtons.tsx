import { Button } from "@mui/joy";
import "../../less/write-action-buttons.less";

type Props = {
  onPublish: () => void;
  onCancel: () => void;
};

function WriteActionButtons({ onPublish, onCancel }: Props) {
  return (
    <div className="write-action-btn-wrapper text-end">
      <Button color="neutral" size="sm" className="text-sm" onClick={onPublish}>
        템플릿 등록
      </Button>
      <Button color="neutral" size="sm" className="text-sm" onClick={onCancel}>
        취소
      </Button>
    </div>
  );
}

export default WriteActionButtons;
