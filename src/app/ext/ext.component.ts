import { Component, OnInit, Input } from '@angular/core';
import { ExtService } from '../api/ext.service';

@Component({
  selector: 'rich-ext',
  templateUrl: './ext.component.html',
  styleUrls: ['./ext.component.scss'],
})
export class ExtComponent implements OnInit {
  @Input()
  type: string
  @Input()
  data: any
  constructor(
    public ext: ExtService
  ) {
  }

  ngOnInit() {
    console.log('type', this.type)
  }

}
