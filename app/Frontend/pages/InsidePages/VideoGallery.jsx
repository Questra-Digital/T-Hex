import React, { useState, useEffect } from "react";
import styles from "./VideoGallery.module.css";
const VideoPlayer = ({ videoUrl }) => {
  return (
    <div className={styles.VideoPlayer}>
      <video controls>
        <source src={videoUrl} type="video/mp4" />
        Your browser does not support the video tag.
      </video>
    </div>
  );
};

const VideoGallery = () => {
  const [selectedOption, setSelectedOption] = useState("");
  const [videoList, setVideoList] = useState([]);
  const [showVideos, setShowVideos] = useState(false);

  const videoMap = {
    video1: ["v1.mp4", "v2.mp4"],
    video2: ["v2.mp4", "v3.mp4"],
  };

  const handleOptionChange = (event) => {
    const selectedOption = event.target.value;
    setSelectedOption(selectedOption);
    setVideoList([]);
  };

  const handleLoadVideos = () => {
    if (selectedOption in videoMap) {
      setVideoList(videoMap[selectedOption]);
      setShowVideos(true);
    } else {
      setVideoList([]);
      setShowVideos(false);
    }
  };
  return (
    <div className="mt-16">
      <h1>Video Gallery</h1>
      <select value={selectedOption} onChange={handleOptionChange}>
        <option value="">Select a video category</option>
        <option value="video1">Video 1</option>
        <option value="video2">Video 2</option>
      </select>
      <button onClick={handleLoadVideos}>Load Videos</button>
      {showVideos && (
        <div className={styles.videoGrid}>
          {videoList.map((videoUrl, index) => (
            <VideoPlayer key={index} videoUrl={videoUrl} />
          ))}
        </div>
      )}
    </div>
  );
};

export default VideoGallery;
