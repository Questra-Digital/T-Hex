import React, { useState } from "react";
import "./payment.module.css";

function PaymentPage() {
  const [cardNumber, setCardNumber] = useState("");
  const [cardholderName, setCardholderName] = useState("");
  const [cvc, setCvc] = useState("");
  const [expirationDate, setExpirationDate] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    // Process payment information here (e.g., send to a payment gateway)
    console.log("Payment details:", {
      cardNumber,
      cardholderName,
      cvc,
      expirationDate,
    });
    // Clear the form after submission
    setCardNumber("");
    setCardholderName("");
    setCvc("");
    setExpirationDate("");
  };

  return (
    <div className="payment-container">
      <div className="payment-card">
        <img
          src="pay.jpg"
          alt="Debit Card"
          className="card-image"
        />
        <div className="card-details" style={{
              marginLeft: "10px",
           
             
              justifyContent: "center",
              backgroundColor: "white",
            }}>
          <label htmlFor="cardNumber">Card Number</label>
          <input
            type="text"
            id="cardNumber"
            value={cardNumber}
            onChange={(e) => setCardNumber(e.target.value)}
            maxLength="16"
            placeholder="**** **** **** ****"
            required
          />
          <label htmlFor="cardholderName" style={{
              marginLeft: "10px",
           
             
              justifyContent: "center",
              backgroundColor: "white",
            }}>Cardholder Name</label>
          <input
            type="text"
            id="cardholderName"
            value={cardholderName}
            onChange={(e) => setCardholderName(e.target.value)}
            required
          />
          <div className="expiration-cvc" style={{
              marginLeft: "10px",
           
             
              justifyContent: "center",
              backgroundColor: "white",
            }}>
            <div className="expiration">
              <label htmlFor="expirationDate">Expiration Date</label>
              <input
                type="text"
                id="expirationDate"
                value={expirationDate}
                onChange={(e) => setExpirationDate(e.target.value)}
                placeholder="MM/YY"
                maxLength="5"
                required
              />
            </div>
            <div className="cvc" style={{
              marginLeft: "10px",
           
             
              justifyContent: "center",
              backgroundColor: "white",
            }}>
              <label htmlFor="cvc">CVC</label>
              <input
                type="password"
                id="cvc"
                value={cvc}
                onChange={(e) => setCvc(e.target.value)}
                maxLength="3"
                placeholder="***"
                required
              />
            </div>
          </div>
        </div>
      </div>
      <form onSubmit={handleSubmit}>
      <a href="/Login" style={{ textDecoration: "none" }}>
  <button
    type="submit"
    className="pay-btn"
    style={{
      backgroundColor: "purple",
      color: "white",
      padding: "10px 20px",
      borderRadius: "5px",
      border: "none",
      cursor: "pointer",
      display: "inline-block",
    }}
  >
    Pay Now and Login
  </button>
</a>

      </form>
    </div>
  );
}

export default PaymentPage;
