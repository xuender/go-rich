import { Component } from '@angular/core';
import { ModalController, ActionSheetController, List } from '@ionic/angular'
import { Observable } from 'rxjs';
import { pull, includes } from 'lodash'

import { CustomerService } from './customer.service';
import { Customer } from './customer';
import { CustomerPage } from './customer/customer.page';
import { UploadPage } from './upload/upload.page';
import { TagService } from '../tags/tag.service';
import { ObjsPage } from '../api/objs.page';

@Component({
  selector: 'app-customers',
  templateUrl: 'customers.page.html',
  styleUrls: ['customers.page.scss']
})
export class CustomersPage extends ObjsPage<Customer>{
  customers$: Observable<Customer[]>
  tags: string[] = []
  constructor(
    public customerService: CustomerService,
    private tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.customers$ = this.customerService.all$
  }

  get service() { return this.customerService }
  get page() { return CustomerPage }
  get newObj() { return { name: '', extend: {}, tags: [] } }
  get title() { return '商品' }

  search(event: CustomEvent) {
    const txt = event.detail.value
    this.customers$ = this.customerService.search(txt)
  }

  async select() {
    const tags = await this.tagService.select('tag-C')
    if (tags) {
      this.tags = tags
      this.customers$ = this.customerService.select(this.tags)
    }
  }

  hasTag(tag: string): boolean {
    return includes(this.tags, tag)
  }

  removeTag(tag: string) {
    pull(this.tags, tag)
    this.customers$ = this.customerService.select(this.tags)
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
