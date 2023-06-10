import { Link } from "react-router-dom";
import "../styles/not-found.less";

function NotFound() {
  return (
    <div className="page-wrapper not-found">
      <div className="page-container">
        <p className="title-text">잘못된 요청입니다</p>
        <div className="action-button-container">
          <Link to="/" className="link-btn">
            시작 페이지로
          </Link>
        </div>
      </div>
    </div>
  );
}

export default NotFound;
