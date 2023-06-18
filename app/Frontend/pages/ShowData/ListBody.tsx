import ListRow from "./ListRow";


const ListBody = ({ data }: any) => {
  
  var id_c=0;

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
        id_c = id_c + 1;
        return (
          <ListRow
            key={idx}
            _id={id_c}
            owner={project.username}
            createdat={project.testfile}
            status={project.status}
            duration={project.duration.substring(0, project.duration.indexOf("."))}
          />
        );
      })}
    </tbody>
  );
};

export default ListBody;
