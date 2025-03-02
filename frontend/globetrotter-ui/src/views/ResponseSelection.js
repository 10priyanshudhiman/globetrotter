import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getRequestToAPI, postRequestToJsonAPI } from "../services/api";
import URLS from "../urls";
import store from "../store/store";
import {userActionResult} from "../store/actions"


const ResponseSelection = () => {
  const { cardId } = useParams();
  const navigate = useNavigate();

  const [clues, setClues] = useState([]);
  const [cities, setCities] = useState([]);
  const [selectedClueId, setSelectedClueId] = useState([]); // Default: all clues selected
  const [selectedCities, setSelectedCities] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchGameDetails = async () => {
      try {
        const response = await getRequestToAPI(`${URLS.gameDetails}?country=${cardId}`);
        if (response.status === 200 && response.data?.data) {
          const { clues, cities } = response.data.data;
          const formattedClues = Object.entries(clues).map(([clueId, clueText]) => ({
            clueId,
            clueText,
          }));
          
          setClues(formattedClues);
          setSelectedClueId(formattedClues.map(clue => clue.clueId)); 
          setCities(cities);
        }
      } catch (error) {
        console.error("Error fetching game details:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchGameDetails();
  }, [cardId]);

  const toggleCitySelection = (city) => {
    setSelectedCities((prevSelected) =>
      prevSelected.includes(city)
        ? prevSelected.filter((c) => c !== city)
        : [...prevSelected, city]
    );
  };

  const submitResponse = async () => {
    if (selectedCities.length === 0) {
      alert("Please select at least one city.");
      return;
    }

    const payload = selectedCities.flatMap(city =>
        selectedClueId.map(clue_id => ({
          country: cardId, // Using cardId as country
          city,
          clue_id
        }))
      );

    try {
      const response = await postRequestToJsonAPI(URLS.verfiyAction, payload);
      if (response.status === 200) {
        console.log("response",response)
        store.dispatch(userActionResult(response?.data?.data) || {})
        navigate("/score");
      }
    } catch (error) {
      console.error("Error verifying response:", error);
    }
  };

  return (
    <div className="container-responses">
      {loading ? (
        <p>Loading...</p>
      ) : (
        <>
          <h2>{cardId}</h2>

          <div className="clues-section">
            <h3>Clues (Pre-selected)</h3>
            {clues.map(({ clueId, clueText }) => (
              <div key={clueId} className="clue-card selected">
                {clueText}
              </div>
            ))}
          </div>

          <div className="cities-section">
            <h3>Select Cities</h3>
            <div className="city-list">
              {cities.map((city, index) => (
                <div
                  key={index}
                  className={`city-card ${selectedCities.includes(city) ? "selected" : ""}`}
                  onClick={() => toggleCitySelection(city)}
                >
                  {city}
                </div>
              ))}
            </div>
          </div>

          <button onClick={submitResponse} disabled={selectedCities.length === 0}>
            Submit Response
          </button>
        </>
      )}
    </div>
  );
};

export default ResponseSelection;
