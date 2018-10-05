import { ToNamePipe } from './to-name.pipe';

describe('ToNamePipe', () => {
  it('create an instance', () => {
    const pipe = new ToNamePipe();
    expect(pipe).toBeTruthy();
  });
});
