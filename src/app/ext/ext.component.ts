import { Component, OnInit, Input } from '@angular/core';

import { Ext } from './ext';
import { ExtService } from './ext.service';

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
  exts: Ext[] = []
  constructor(
    private extService: ExtService
  ) {
  }

  ngOnInit() {
    if (this.type === 'C') {
      this.extService.extsByCustomer$.subscribe((exts) => this.exts = exts)
    } else {
      this.extService.extsByItem$.subscribe((exts) => this.exts = exts)
    }
  }

  async edit() {
    if (this.type === 'C') {
      await this.extService.extCustomer()
      this.exts = await this.extService.extsByCustomer$.toPromise()
    } else {
      await this.extService.extItem()
      this.exts = await this.extService.extsByItem$.toPromise()
    }
  }
}
