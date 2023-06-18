import React from "react";

const Styles = {
  contactContainer: {
    marginTop: "80px",
    marginBottom: "100px",
    textAlign: "center",
  },

  contactForm: {
    maxWidth: "500px",
    margin: "auto",
  },

  contactInputs: {
    display: "flex",
    flexDirection: "column",
    gap: "30px",
  },

  inputField: {
    width: "100%",
    padding: "12px 20px",
    margin: "8px 0",
    display: "inline-block",
    border: "1px solid #ccc",
    borderRadius: "4px",
    boxSizing: "border-box",
  },

  contactInputSubmit: {
    width: "100%",
    backgroundColor: "#1a83ff",
    color: "white",
    padding: "14px 20px",
    margin: "8px 0",
    border: "none",
    borderRadius: "4px",
    cursor: "pointer",
  },

  contactInputSubmit: {
    ":hover": {
      backgroundColor: "#ff014d",
    },
  },
  textarea: {
    width: "100%",
    height: "150px",
    padding: "12px 20px",
    boxSizing: "border-box",
    border: "2px solid #ccc",
    borderRadius: "4px",
    backgroundColor: "#f8f8f8",
    fontSize: "16px",
    resize: "none",
  },
};

function ContactUs() {
  return (
    <div
      style={{
        marginTop: "50px",
        marginRight: "30px",
        marginBottom: "200px",
        textAlign: "center",
      }}
    >
      <div style={{ margin: "0px", textAlign: "center" }}>
        <h1
          style={{
            color: "Black",
            textTransform: "capitalize",
            marginBottom: "30px",
            fontSize: "30px",
          }}
        >
          Feel Free to Contact us
        </h1>
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d13610.206048358172!2d74.3030141!3d31.4815212!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x391903f08ebc7e8b%3A0x47e934f4cd34790!2sFAST%20NUCES%20Lahore!5e0!3m2!1sen!2s!4v1681610757653!5m2!1sen!2s"
          width="600"
          height="450"
          style={{ border: "0", display: "block", margin: "0 auto" }}
          allowFullScreen=""
          loading="lazy"
          referrerpolicy="no-referrer-when-downgrade"
        ></iframe>
      </div>

      <div style={Styles.contactContainer} className="contact-container">
        <div style={Styles.contactForm} className="contact-form">
          <form
            action="#"
            method="POST"
            style={Styles.contactInputs}
            className="contact-inputs"
          >
            <input
              className="input-field"
              style={Styles.inputField}
              type="text"
              name="username"
              placeholder="Username"
              autoComplete="off"
              required
            />

            <input
              className="input-field"
              type="email"
              name="Email"
              placeholder="Email"
              autoComplete="off"
              required
              style={Styles.inputField}
            />

            <textarea
              name="message"
              cols="30"
              rows="6"
              placeholder="Write message"
              autoComplete="off"
              required
              style={Styles.textarea}
            ></textarea>

            <input style={Styles.inputField} type="submit" value="SEND" />
          </form>
        </div>
      </div>
    </div>
  );
}
export default ContactUs;
