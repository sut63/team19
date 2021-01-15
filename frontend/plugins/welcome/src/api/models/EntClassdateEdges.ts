/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCourseclass,
    EntCourseclassFromJSON,
    EntCourseclassFromJSONTyped,
    EntCourseclassToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClassdateEdges
 */
export interface EntClassdateEdges {
    /**
     * Courseclasses holds the value of the courseclasses edge.
     * @type {Array<EntCourseclass>}
     * @memberof EntClassdateEdges
     */
    courseclasses?: Array<EntCourseclass>;
}

export function EntClassdateEdgesFromJSON(json: any): EntClassdateEdges {
    return EntClassdateEdgesFromJSONTyped(json, false);
}

export function EntClassdateEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClassdateEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'courseclasses': !exists(json, 'courseclasses') ? undefined : ((json['courseclasses'] as Array<any>).map(EntCourseclassFromJSON)),
    };
}

export function EntClassdateEdgesToJSON(value?: EntClassdateEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'courseclasses': value.courseclasses === undefined ? undefined : ((value.courseclasses as Array<any>).map(EntCourseclassToJSON)),
    };
}


