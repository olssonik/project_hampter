import './App.css';
import React, { useRef, useEffect } from 'react';

function VideoComponent() {
  const videoRef = useRef(null);

  useEffect(() => {
    const videoElement = videoRef.current;

    const fetchAndSetVideoSource = async () => {
      try {
        const response = await fetch('http://127.0.0.1:8080/video');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        videoElement.src = URL.createObjectURL(await response.blob());
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
      }
    };

    fetchAndSetVideoSource();
  }, []);

  return (
    <video ref={videoRef} controls style={{
      maxWidth: '640px',
      maxHeight: '360px',
      width: '100%',
      height: 'auto'
    }}>
      Your browser does not support the video tag.
    </video >
  );
}


function App() {
  return (
    <>
      <div style={{ maxWidth: '800px', margin: 'auto' }}>
        <VideoComponent />
      </div>
    </>
  );
}

export default App;
