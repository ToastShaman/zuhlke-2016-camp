import { Gender, Person } from './person';

function bob(getAge: Function): Person  {
  return {
    firstname: 'Bob',
    lastname: 'Hope',
    age: getAge(),
    gender: Gender.MALE
  };
}

const aPerson: Person = bob(() => 42);

console.log(`
  I created a person:
    firstname: ${aPerson.firstname}
    lastname: ${aPerson.lastname}
    age:  ${aPerson.age}
    gender:  ${aPerson.gender}
`);
