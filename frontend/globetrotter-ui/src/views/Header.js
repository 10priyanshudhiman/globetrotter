import React, { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { connect } from "react-redux";
import { Modal, Button } from "antd";
import store from "../store/store";
import { updateUser } from "../store/actions";

const Header = ({ user }) => {
  console.log("user", user);
  const userId = user?.userId || "Guest";
  const totalScore = user?.totalScore || 0;
  const correctAnswers = user?.correctAnswers || 0;
  const incorrectAnswers = user?.incorrectAnswers || 0;
  const navigate = useNavigate();
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [imageUrl, setImageUrl] = useState(null);
  const canvasRef = useRef(null);

  const handleLogout = () => {
    localStorage.removeItem("userId");
    store.dispatch(updateUser({}));
    navigate("/");
  };

  const inviteUrl = `${window.location.origin}`;
  const whatsappShareUrl =  `https://api.whatsapp.com/send?text=Join me in this game! My score is ${totalScore}. Play now: ${inviteUrl}`;

  useEffect(() => {
    if (isModalVisible) {
     
    generateInviteImage(totalScore,correctAnswers,incorrectAnswers);
     
    }
  }, [isModalVisible,totalScore,correctAnswers,incorrectAnswers]);

  const generateInviteImage = (totalScore,correctAnswers,incorrectAnswers) => {
    
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");

    // Set canvas size
    canvas.width = 500;
    canvas.height = 300;

    // Constant scores
    

    // Generate random colors for gradient
    const getRandomColor = () => `#${Math.floor(Math.random() * 16777215).toString(16).padStart(6, '0')}`;
    const color1 = getRandomColor();
    const color2 = getRandomColor();

    // Background gradient
    const gradient = ctx.createLinearGradient(0, 0, 500, 300);
    gradient.addColorStop(0, color1);
    gradient.addColorStop(1, color2);
    ctx.fillStyle = gradient;
    ctx.fillRect(0, 0, 500, 300);

    // Title
    ctx.fillStyle = "#fff";
    ctx.font = "bold 24px Arial";
    ctx.fillText("🎮 Join My Game! 🎮", 130, 50);

    // Score section
    ctx.fillStyle = "#fff";
    ctx.font = "20px Arial";
    ctx.fillText(`🏆 Score: ${totalScore}`, 170, 100);
    ctx.fillText(`✅ Correct: ${correctAnswers}`, 170, 140);
    ctx.fillText(`❌ Incorrect: ${incorrectAnswers}`, 170, 180);

    // Invite text
    ctx.font = "18px Arial";
    ctx.fillText("Click the link below to challenge me!", 100, 230);

    // Convert to image
    const imgData = canvas.toDataURL("image/png");
    setImageUrl(imgData);
};


  return (
    <div className="header">
      <span><strong>User:</strong> {userId}</span>
      <span><strong>Total Score:</strong> {totalScore}</span>
      <span><strong>Correct Answers:</strong> {correctAnswers}</span>
      <span><strong>Incorrect Answers:</strong> {incorrectAnswers}</span>
      <button onClick={handleLogout}>Logout</button>
      <button onClick={() => setIsModalVisible(true)} className="invite-button">Invite a Friend</button>


      <Modal title="Invite a Friend" open={isModalVisible} onCancel={() => setIsModalVisible(false)} footer={null}>
        <p>Share this link with your friends to challenge them!</p>
        <input type="text" value={inviteUrl} readOnly className="invite-link" />

        <canvas ref={canvasRef} style={{ display: "none" }}></canvas>

        {imageUrl && <img src={imageUrl} alt="Invite Preview" style={{ width: "100%", marginTop: "10px" }} />}
        {console.log("imageUrl",whatsappShareUrl)}
        <Button type="primary" href={whatsappShareUrl} target="_blank" style={{ marginTop: "10px" }}>
          Share on WhatsApp
        </Button>
      </Modal>
    </div>
  );
};

const mapStateToProps = ({ data: { user } }) => ({
  user,
});

export default connect(mapStateToProps)(Header);
