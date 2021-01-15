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
    EntCourseEdges,
    EntCourseEdgesFromJSON,
    EntCourseEdgesFromJSONTyped,
    EntCourseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCourse
 */
export interface EntCourse {
    /**
     * CourseName holds the value of the "Course_name" field.
     * @type {string}
     * @memberof EntCourse
     */
    courseName?: string;
    /**
     * 
     * @type {number}
     * @memberof EntCourse
     */
    degreeId?: number;
    /**
     * 
     * @type {EntCourseEdges}
     * @memberof EntCourse
     */
    edges?: EntCourseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCourse
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntCourse
     */
    instructorInfoId?: number;
    /**
     * 
     * @type {number}
     * @memberof EntCourse
     */
    subjectId?: number;
}

export function EntCourseFromJSON(json: any): EntCourse {
    return EntCourseFromJSONTyped(json, false);
}

export function EntCourseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCourse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'courseName': !exists(json, 'Course_name') ? undefined : json['Course_name'],
        'degreeId': !exists(json, 'degree_id') ? undefined : json['degree_id'],
        'edges': !exists(json, 'edges') ? undefined : EntCourseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'instructorInfoId': !exists(json, 'instructorInfo_id') ? undefined : json['instructorInfo_id'],
        'subjectId': !exists(json, 'subject_id') ? undefined : json['subject_id'],
    };
}

export function EntCourseToJSON(value?: EntCourse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Course_name': value.courseName,
        'degree_id': value.degreeId,
        'edges': EntCourseEdgesToJSON(value.edges),
        'id': value.id,
        'instructorInfo_id': value.instructorInfoId,
        'subject_id': value.subjectId,
    };
}


