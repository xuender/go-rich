import { Pipe, PipeTransform } from '@angular/core'
@Pipe({ name: 'concat' })
export class ConcatPipe implements PipeTransform {
  transform(value: any[], other: any[] = []): any[] {
    if (value) {
      other.push(...value)
    }
    return other
  }
}
