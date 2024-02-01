import ListRow from "./ListRow";


const ListBody = ({ data }: any) => {
  
  var id_c=0;

  if (!data) {
    return null;
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
