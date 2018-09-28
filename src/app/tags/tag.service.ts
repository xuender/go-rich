import { ModalController } from '@ionic/angular';
import { Injectable, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { find } from 'lodash'

import { Tag } from './tag';
import { Key } from '../api/key';
import { URL } from '../api/init'
import { SelectPage } from '../tags/select/select.page';
import { ObjService } from '../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class TagService extends ObjService<Tag>{
  keys: Key[] = []
  constructor(
    private modalCtrl: ModalController,
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
    this.keys.push(
      { key: 'tag-C', name: '客户', color: 'medium', tags: [], },
      { key: 'tag-I', name: '商品', color: 'secondary', tags: [], },
    )
  }

  path() { return '/api/tags' }

  tagsByKey(key: string, all = true) {
    return this.http.get<Tag[]>(`${URL}/api/tags/${key}?all=${all}`)
  }

  async selectTags(key: string) {
    const k = find(this.keys, k => k.key == key)
    const modal = await this.modalCtrl.create({
      component: SelectPage,
      componentProps: {
        tags: k.tags,
        tags$: this.tagsByKey(key),
      }
    });
    await modal.present()
    const d = await modal.onDidDismiss()
    if (d.data) {
      k.tags = d.data
      return d.data
    }
    return null
  }
}
