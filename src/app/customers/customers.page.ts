import { Component } from '@angular/core';
import { ModalController, ActionSheetController, List, InfiniteScroll } from '@ionic/angular'
import { Observable } from 'rxjs';
import { pull, includes } from 'lodash'

import { CustomerService } from './customer.service';
import { Customer } from './customer';
import { CustomerPage } from './customer/customer.page';
import { UploadPage } from './upload/upload.page';
import { TagService } from '../tags/tag.service';
import { ObjsPage } from '../api/objs.page';
import { Paging } from '../api/paging';

@Component({
  selector: 'app-customers',
  templateUrl: 'customers.page.html',
  styleUrls: ['customers.page.scss']
})
export class CustomersPage extends ObjsPage<Customer>{
  tags: string[] = []
  paging: Paging<Customer> = { data: [], total: 0 }
  constructor(
    public customerService: CustomerService,
    private tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.customerService.paging$.subscribe(p => this.paging = p)
  }

  get service() { return this.customerService }
  get page() { return CustomerPage }
  get newObj() { return { name: '', extend: {}, tags: [] } }
  get title() { return '客户' }

  search(event: CustomEvent) {
    const txt = event.detail.value
    this.customerService.searchPaging(txt).subscribe(p => this.paging = p)
  }

  async select() {
    const tags = await this.tagService.select('tag-C')
    if (tags) {
      this.tags = tags
      this.customerService.selectPaging(this.tags).subscribe(p => this.paging = p)
    }
  }

  load(event) {
    this.customerService.nextPaging$.subscribe(p => {
      console.log('p', p)
      this.paging.data.push(...p.data)
      event.target.complete()
      event.target.disabled = p.total == p.data.length
    })
  }

  hasTag(tag: string): boolean {
    return includes(this.tags, tag)
  }

  removeTag(tag: string) {
    pull(this.tags, tag)
    this.customerService.paging$.subscribe(p => this.paging = p)
  }

  // 上传客户xslx
  async upload() {
    const modal = await this.modalCtrl.create({
      component: UploadPage,
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
      }
    })
    return await modal.present()
  }
}
