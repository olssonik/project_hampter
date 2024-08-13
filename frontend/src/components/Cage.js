import { useEffect, useState } from "react";

function Cage() {
  const [cage, setCage] = useState([])

  useEffect(() => {
    setCage(Array(9).fill(0))
    console.log(cage)
  }, [cage])
  return (
    <div>

      {cage}

    </div>
  );
}

export default Cage;
