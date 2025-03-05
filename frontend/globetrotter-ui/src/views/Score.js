import React, { useEffect, useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { connect } from "react-redux";
import Confetti from "react-confetti";
import { postRequestToJsonAPI } from "../services/api";
import URLS from "../urls";
import store from "../store/store";
import { updateUser } from "../store/actions";

const Score = ({ userActionResult, user }) => {
  console.log("action result", userActionResult, user);
  const navigate = useNavigate();
  const hasUpdated = useRef(false); 

  const [showConfetti, setShowConfetti] = useState(false);
  const [showSadFace, setShowSadFace] = useState(false);
  const [funFact, setFunFact] = useState("");

  useEffect(() => {
    console.log("inside", userActionResult?.isCorrect);

    if (userActionResult?.isCorrect !== undefined && !hasUpdated.current && user?.userId) {
      hasUpdated.current = true; 

      const { isCorrect, facts } = userActionResult;
      const point = isCorrect ? 1 : -1;
      const userId = user?.userId;

      const payload = {
        userId: userId,
        point: point,
        isCorrect: isCorrect,
      };

      postRequestToJsonAPI(URLS.updateUser, payload)
        .then((res) => {
          console.log("data", res.data.data);
          store.dispatch(updateUser(res?.data?.data))
          hasUpdated.current = false
        })
        .catch((err) => {
          console.log("error", err);
          hasUpdated.current = false
        });

      if (isCorrect) {
        setShowConfetti(true);
      } else {
        setShowSadFace(true);
      }

      const factKeys = Object.keys(facts || {});
      if (factKeys.length > 0) {
        setFunFact(facts[factKeys[Math.floor(Math.random() * factKeys.length)]]);
      }
    }
  }, [userActionResult,user?.userId]);

  return (
    <div className="container" style={{ textAlign: "center", padding: "20px" }}>
      {showConfetti && <Confetti />}
      {showSadFace && <p style={{ fontSize: "40px" }}>😢</p>}

      <h2>Your Score</h2>
      <p style={{ fontSize: "32px", fontWeight: "bold", color: "#007bff" }}>{userActionResult?.isCorrect ? 1 : -1}</p>

      {funFact && (
        <div style={{ marginTop: "20px", fontSize: "18px", fontStyle: "italic" }}>
          Fun Fact: {funFact}
        </div>
      )}

      <button onClick={() => navigate("/dashboard")} style={{ marginTop: "20px" }}>
        Play Again
      </button>
    </div>
  );
};

const mapStateToProps = ({ data: { userActionResult, user } }) => ({
  userActionResult,
  user,
});

export default connect(mapStateToProps)(Score);
