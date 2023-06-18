import { useState } from "react";
import ListRow from "./ListRow";


const ListBody = ({ data }: any) => {
 
  if (!data) {
    // Handle the case when data is undefined or null
    return null; // or render an appropriate fallback component or message
  }
  else{
    console.log(data);
  }


  return (
    <tbody>
      {data.map((project: any, idx: any) => {
        return (
          <ListRow
            key={idx}
            createdat={project.testfile}
            path={project.urlid}
          />
        );
      })}
    </tbody>
  );
};

export default ListBody;
