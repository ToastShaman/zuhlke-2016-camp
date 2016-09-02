import { Animal } from './animal';
import { Dog } from './dog';
import { Smell } from './smell';

// new Animal(); // Cannot create an instance of the abstract class 'Animal'.

const woof = new Dog();
woof.say('Hello');

const smell = woof.doSmell();
console.log(smell.description);
