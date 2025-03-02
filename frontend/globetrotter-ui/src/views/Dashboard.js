import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getRequestToAPI } from "../services/api";
import URLS from "../urls";

const Dashboard = () => {
  const navigate = useNavigate();
  const [countries, setCountries] = useState([]);

  useEffect(() => {
    const fetchCountries = async () => {
      try {
        const response = await getRequestToAPI(URLS.getAllCountries);
        if (response.status === 200 && response.data) {
          const countryData = response.data.data; // Extract 'data' field
          
          if (typeof countryData === "object" && countryData !== null) {
            const countryList = Object.keys(countryData); // Get only country names
            setCountries(countryList);
          } else {
            console.error("Invalid data format", countryData);
          }
        }
      } catch (error) {
        console.error("Error fetching countries:", error);
      }
    };

    fetchCountries();
  }, []);

  return (
    <div>
      <h1 style={{ textAlign: "center", marginTop: "20px" }}>Select a Country</h1>
      <div className="cards-container">
        {countries.length > 0 ? (
          countries.map((country, index) => (
            <div 
              key={index} 
              className="card" 
              onClick={() => navigate(`/responses/${country}`)}
            >
              {country}
            </div>
          ))
        ) : (
          <p style={{ textAlign: "center" }}>Loading countries...</p>
        )}
      </div>
    </div>
  );
};

export default Dashboard;
