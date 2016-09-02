import { Animal } from './animal';
import { Nose, Smell } from './smell';

export class Dog extends Animal implements Nose {
  
  public say(message: string): void {
    console.log(`I am a dog saying: ${message}`);
  }

  public doSmell(): Smell {
    return { description: 'smells like flowers...' } as Smell;
  }
}
